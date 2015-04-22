define(function(require) {

  var Backbone = require('backbone');
  var Tile = require('app/models/tile');

  var Board = Backbone.Collection.extend({

    model: Tile,

    dimension: function() {
      return Math.sqrt(this.length);
    },

    swap: function(a, b) {
      var bIndex = b.index();
      this.models[a.index()] = b;
      this.models[bIndex] = a;
      a.trigger('swap');
      b.trigger('swap');
    },

    move: function(direction) {
      var blank = this.findWhere({ value: Board.BLANK });

      switch (direction) {
        case Board.UP:
          target = this.at(blank.index() - this.dimension());
          break;
        case Board.DOWN:
          target = this.at(blank.index() + this.dimension());
          break;
        case Board.LEFT:
          target = this.at(blank.index() - 1);
          break;
        case Board.RIGHT:
          target = this.at(blank.index() + 1);
          break;
        default:
          return;
      }

      this.swap(target, blank);
    }

  }, {

    BLANK: 8,

    UP: 'Up',
    DOWN: 'Down',
    LEFT: 'Left',
    RIGHT: 'Right'

  });

  return Board;

});
