define(function(require) {

  var _ = require('underscore');
  var Backbone = require('backbone');
  var Leaderboard = require('app/collections/leaderboard');
  var Handlebars = require('handlebars');
  var LeaderboardEntryView = require('app/views/leaderboard_entry');

  return Backbone.View.extend({

    template: Handlebars.compile(require('text!app/templates/leaderboard.hbs')),

    className: 'leaderboard',

    initialize: function() {
      this.collection = new Leaderboard();
      this.collection.fetch();
      this.collection.on('add', this.addEntry, this);

      $(window).on('scroll', _.debounce(this.onScroll.bind(this), 150));
    },

    onScroll: function() {
      if (this.shouldLoadMore()) {
        this.collection.next();
      }
    },

    shouldLoadMore: function() {
      var scrollY = window.scrollY || document.documentElement.scrollTop;
      return (window.innerHeight + scrollY) >= (document.body.offsetHeight - 400);
    },

    addEntry: function(entryModel) {
      var entryView = new LeaderboardEntryView({ model: entryModel });
      entryView.render();
      this.$('.js-entries').append(entryView.el);
    },

    render: function() {
      this.$el.html(this.template());
    }

  });

});
