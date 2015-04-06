define(function(require) {

  var Backbone = require('backbone');
  var Board = require('app/collections/board');

  return Backbone.Model.extend({

    urlRoot: '/api/v1/puzzles',

    parse: function(attributes) {
      attributes.board = new Board(attributes.board.state.map(function(tile) {
        return { value: tile };
      }));

      return attributes;
    }

  });

});
