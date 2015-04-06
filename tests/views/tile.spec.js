define(function(require) {

  var _ = require('underscore');
  var Board = require('app/collections/board');
  var Tile = require('app/views/tile');

  describe('app/views/tile', function() {

    var board;

    beforeEach(function() {
      board = new Board(_.range(9).map(function(value) {
        return { value: value };
      }));
    });

    it('moves when clicked and adjacent to the blank', function() {
      var tile = new Tile({
        model: board.at(1),
        collection: board
      });
      tile.click();
      expect(board.at(0).get('value')).toEqual(tile.model.get('value'));
    });

    it('does not move when clicked and not adjacent to the blank', function() {
      var tile = new Tile({
        model: board.at(8),
        collection: board
      });
      tile.click();
      expect(board.at(8).get('value')).toEqual(tile.model.get('value'));
    });

  });

});
