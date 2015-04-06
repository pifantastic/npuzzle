define(function(require) {

  var Backbone = require('backbone');

  return Backbone.Model.extend({

    urlRoot: '/api/v1/solutions',

    defaults: function() {
      return {
        moves: []
      };
    },

    fetch: function() {
      var body = {
        board: this.get('board').map(function(tile) {
          return tile.get('value');
        })
      };

      var request = $.post(this.urlRoot, JSON.stringify(body));

      request.done(function(moves) {
        this.set('moves', moves);
      }.bind(this));

      return request;
    }

  });

});
