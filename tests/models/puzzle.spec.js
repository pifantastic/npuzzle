define(function(require) {

  var _ = require('underscore');
  var Board = require('app/collections/board');
  var Puzzle = require('app/models/puzzle');

  describe('app/models/puzzle', function() {

    beforeEach(function() {
      jasmine.Ajax.install();
    });

    afterEach(function() {
      jasmine.Ajax.uninstall();
    });

    it('generates the correct API request', function() {
      var puzzle = new Puzzle();
      puzzle.fetch();

      var request = jasmine.Ajax.requests.mostRecent();
      expect(request.url).toBe('/api/v1/puzzles');
      expect(request.method).toBe('GET');

      jasmine.Ajax.requests.mostRecent().respondWith({
        status: 200,
        contentType: 'application/json',
        responseText: '{"dimension":3,"board":{"state":[0,1,2,3,4,5,6,7,8]}}'
      });

      expect(puzzle.get('dimension')).toEqual(3);
      expect(puzzle.get('board').dimension()).toEqual(3);
    });

  });

});
