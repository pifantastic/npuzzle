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
      this.shuffle();
    },

    shuffle: function() {
      this.model.fetch();
    },

    theme: function(e) {
      var theme = this.$(e.currentTarget).val();
      this.model.get('board').invoke('set', 'theme', theme);
    },

    solve: function() {
      var solution = new Solution({
        board: this.model.get('board')
      });

      solution.once('change', function() {
        var moves = solution.get('moves');

        setInterval(function() {
          this.model.get('board').move(moves.shift());
        }.bind(this), 400);
      }.bind(this));

      solution.fetch();
    },

    render: function() {
      var $tiles = this.$('.tiles');

      var width = $tiles.width() / this.model.get('dimension');
      var height = $tiles.height() / this.model.get('dimension');

      $tiles.empty();

      this.model.get('board').forEach(function(tile) {
        tile.set('width', width);
        tile.set('height', height);

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
