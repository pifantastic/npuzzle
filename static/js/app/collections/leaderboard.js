define(function(require) {

  var Backbone = require('backbone');
  var LeaderboardEntry = require('app/models/leaderboard_entry');

  var Leaderboard = Backbone.Collection.extend({

    model: LeaderboardEntry,

    initialize: function() {
      this.page = 1;
    },

    url: function() {
      return '/api/v1/leaderboards?page=' + this.page;
    },

    next: function() {
      this.page += 1;
      this.fetch({ remove: false });
    },

    parse: function(response) {
      return response.entries;
    }

  });

  return Leaderboard;

});
