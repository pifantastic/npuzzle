define(function(require) {

  var _ = require('underscore');
  var Board = require('app/collections/board');

  describe('app/models/puzzle', function() {

    var board;

    beforeEach(function() {
      board = new Board(_.range(9).map(function(value) {
        return { value: value };
      }));
    });

    it('calculates the correct dimension', function() {
      expect(board.dimension()).toEqual(3);
    });


    it('swaps tiles correctly', function() {
      var first = board.at(0);
      var second = board.at(1);

      var firstSwap = jasmine.createSpy('first');
      first.on('swap', firstSwap);

      var secondSwap = jasmine.createSpy('second');
      second.on('swap', secondSwap);

      board.swap(first, second);

      expect(board.at(0).get('value')).toEqual(1);
      expect(board.at(1).get('value')).toEqual(0);

      expect(firstSwap).toHaveBeenCalled();
      expect(secondSwap).toHaveBeenCalled();
    });

    it('moves tiles correctly', function() {
      board.move(Board.RIGHT);
      expect(board.at(1).get('value')).toEqual(0);
    });

  });

});
