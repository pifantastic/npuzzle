define(function(require) {

  var $ = require('jquery');
  var Backbone = require('backbone');

  var PuzzleView = require('app/views/puzzle');
  var LeaderboardView = require('app/views/leaderboard');

  var Router = Backbone.Router.extend({

    routes: {
      "":                   "puzzle",
      "leaderboard":        "leaderboard",
      "leaderboard/p:page": "leaderboard"
    },

    puzzle: function() {
      $('.js-nav-home').addClass('active');

      new PuzzleView({
        el: '.js-puzzle'
      });
    },

    leaderboard: function(page) {
      $('.js-nav-leaderboard').addClass('active');

      new LeaderboardView({
        el: '.js-leaderboard'
      }).render();
    }

  });

  new Router();

  Backbone.history.start({
    pushState: true
  });

});
