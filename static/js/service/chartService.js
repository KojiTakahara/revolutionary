var service = angular.module('chartService', []);
service.factory('chartService', ['$http', '$q', function($http, $q) {
  var service = {};

  service.createTypeColumns = function(data) {
    var obj = {};
    for (var i = 0; i < data.length; i++) {
      if (data[i].Type === "") {
        continue;
      }
      if (obj[data[i].Type] === undefined) {
        obj[data[i].Type] = 1;
      } else {
        obj[data[i].Type]++;
      }
    }
    var columns = [];
    for (var key in obj) {
      columns.push([key, obj[key]]);
    }
    return columns;
  };

  service.createRaceColumns = function(data) {
    var obj = {};
    for (var i = 0; i < data.length; i++) {
      if (data[i].Race === "") {
        continue;
      }
      if (obj[data[i].Race] === undefined) {
        obj[data[i].Race] = 1;
      } else {
        obj[data[i].Race]++;
      }
    }
    var columns = [];
    for (var key in obj) {
      columns.push([key, obj[key]]);
    }
    return columns;
  };

  service.createColorColumns = function(data) {
    var light = ["light"],
        water = ["water"],
        dark = ["dark"],
        fire = ["fire"],
        nature = ["nature"],
        zero = ["zero"];
    for (var i = 0; i < data.length; i++) {
      if (data[i].Light) {
        light.push(1);
      }
      if (data[i].Water) {
        water.push(1);
      }
      if (data[i].Dark) {
        dark.push(1);
      }
      if (data[i].Fire) {
        fire.push(1);
      }
      if (data[i].Nature) {
        nature.push(1);
      }
      if (data[i].Zero) {
        zero.push(1);
      }
    }
    return [light, water, dark, fire, nature, zero];
  };

  return service;
}]);
