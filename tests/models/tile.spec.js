define(function(require) {

  var _ = require('underscore');
  var Board = require('app/collections/board');

  describe('app/models/tile', function() {

    var board, first, last;

    beforeEach(function() {
      board = new Board(_.range(9).map(function(value) {
        return { value: value };
      }));

      first = board.at(0);
      last = board.at(8);
    });

    it('has the correct value', function() {
      expect(first.get('value')).toEqual(0);
      expect(last.get('value')).toEqual(8);
    });

    it('has the correct index', function() {
      expect(first.index()).toEqual(0);
      expect(last.index()).toEqual(8);
    });

    it('has the correct row', function() {
      expect(first.row()).toEqual(0);
      expect(last.row()).toEqual(2);
    });

    it('has the correct column', function() {
      expect(first.col()).toEqual(0);
      expect(last.col()).toEqual(2);
    });

    it('has the correct position', function() {
      expect(first.position()).toEqual({ row: 0, col: 0 });
      expect(last.position()).toEqual({ row: 2, col: 2 });
    });

    it('calculates distance correctly', function() {
      expect(first.distance(last)).toEqual(4);
    });

  });

});
