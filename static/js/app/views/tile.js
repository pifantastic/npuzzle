define(function(require) {

  var Backbone = require('backbone');
  var Board = require('app/collections/board');

  return Backbone.View.extend({

    className: 'tile',

    events: {
      'click': 'click'
    },

    initialize: function() {
      this.model.on('swap', this.move, this);
      this.model.on('change:theme', this.render, this);
    },

    click: function(e) {
      var blank = this.collection.findWhere({
        value: Board.BLANK
      });

      if (this.model.distance(blank) === 1) {
        this.collection.swap(this.model, blank);
      }
    },

    move: function() {
      this.$el.css({
        left: this.model.col() * this.model.get('width'),
        top: this.model.row() * this.model.get('height'),
      });
    },

    render: function() {
      var left = this.model.col() * this.model.get('width');
      var top = this.model.row() * this.model.get('height');

      var bgLeft = this.model.get('value') % this.collection.dimension();
      var bgTop = Math.floor(this.model.get('value') / this.collection.dimension());

      this.$el.css({
        width: this.model.get('width'),
        height: this.model.get('height'),
        left: left,
        top: top,
        'background-position-x': -(bgLeft * this.model.get('width')) + 'px',
        'background-position-y': -(bgTop * this.model.get('height')) + 'px'
      });

      this.$el
        .attr('class', this.className)
        .addClass(this.model.get('theme'))
        .toggleClass('blank', this.model.get('value') === Board.BLANK)
        .data('tile', this.model.get('value'));

      return this;
    }

  });

});
