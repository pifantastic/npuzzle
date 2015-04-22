define(function(require) {

  var Backbone = require('backbone');
  var Handlebars = require('handlebars');

  return Backbone.View.extend({

    tagName: 'li',

    className: 'entry',

    template: Handlebars.compile(require('text!app/templates/leaderboard_entry.hbs')),

    render: function() {
      this.$el.html(this.template(this.model.attributes));

      return this;
    }

  });

});
