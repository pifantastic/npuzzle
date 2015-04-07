define(function(require) {

  var Backbone = require('backbone');
  var Puzzle = require('app/models/puzzle');
  var Solution = require('app/models/solution');
  var TileView = require('app/views/tile');

  return Backbone.View.extend({

    events: {
      'click .js-solve': 'solve',
      'click .js-shuffle': 'shuffle',
      'change .js-theme': 'theme'
    },

    initialize: function() {
      this.model = new Puzzle();
      this.model.on('sync', this.render, this);
      this.model.on('change:solving', this.solving, this);
      this.shuffle();
    },

    shuffle: function() {
      this.model.fetch();
    },

    theme: function(e) {
      var theme = this.$(e.currentTarget).val();
      this.model.get('board').invoke('set', 'theme', theme);
    },

    solving: function() {
      this.$el.toggleClass('solving', this.model.get('solving'));
      this.$('.js-solve').text(this.model.get('solving') ? 'Solving' : 'Solve');
      this.$('button, select').prop('disabled', this.model.get('solving'));
    },

    solve: function() {
      if (this.model.get('solving')) {
        return true;
      }

      this.model.set('solving', true);

      var solution = new Solution({
        board: this.model.get('board')
      });

      solution.once('change', function() {
        var moves = solution.get('moves');
        this.model.get('board').move(moves.shift());

        var interval = setInterval(function() {
          this.model.get('board').move(moves.shift());

          if (moves.length === 0) {
            this.model.set('solving', false);
            clearInterval(interval);
          }
        }.bind(this), 300);
      }.bind(this));

      solution.fetch();
    },

    render: function() {
      this.$el.addClass('puzzle-' + this.model.get('dimension'));
      var $tiles = this.$('.js-tiles');

      var width = $tiles.width() / this.model.get('dimension');
      var height = $tiles.height() / this.model.get('dimension');
      var theme = this.$('.js-theme').val();

      $tiles.empty();

      this.model.get('board').forEach(function(tile) {
        tile.set('width', width);
        tile.set('height', height);
        tile.set('theme', theme);

        var view = new TileView({
          model: tile,
          collection: this.model.get('board')
        }).render();

        $tiles.append(view.el);
      }, this);

      return this;
    }

  });

});
