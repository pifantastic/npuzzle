define(function(require) {

  var _ = require('underscore');
  var Board = require('app/collections/board');
  var Solution = require('app/models/solution');

  describe('app/models/solution', function() {

    var board;

    beforeEach(function() {
      board = new Board(_.range(9).map(function(value) {
        return { value: value };
      }));

      board.move(Board.RIGHT);

      jasmine.Ajax.install();
    });

    afterEach(function() {
      jasmine.Ajax.uninstall();
    });

    it('generates the correct API request', function() {
      var solution = new Solution({ board: board });
      solution.fetch();

      var request = jasmine.Ajax.requests.mostRecent();
      expect(request.url).toBe('/api/v1/solutions');
      expect(request.method).toBe('POST');
      expect(request.params).toEqual(JSON.stringify({ board: [1, 0, 2, 3, 4, 5, 6, 7, 8] }));

      jasmine.Ajax.requests.mostRecent().respondWith({
        status: 200,
        contentType: 'application/json',
        responseText: '["LEFT"]'
      });

      expect(solution.get('moves')).toEqual(['LEFT']);
    });

  });

});
