define(function(require) {

  var Backbone = require('backbone');

  return Backbone.Model.extend({

    defaults: {
      value: 0,
      blank: false,
      theme: 'cat',
      width: 200,
      height: 200
    },

    index: function() {
      return this.collection.indexOf(this);
    },

    row: function() {
      return Math.floor(this.index() / this.collection.dimension());
    },

    col: function() {
      return this.index() % this.collection.dimension();
    },

    position: function() {
      return {
        row: this.row(),
        col: this.col()
      };
    },

    distance: function(other) {
      return Math.abs(this.row() - other.row()) + Math.abs(this.col() - other.col());
    }

  });

});
