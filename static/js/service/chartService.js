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

  service.createTime = function(data) {
    var obj = {};
    var x = ['x'];
    for (var i = 0; i < data.length; i++) {
      if (data[i].Type === "") {
        continue;
      }
      if (obj[data[i].Date] === undefined) {
        obj[data[i].Date] = [data[i]];
      } else {
        obj[data[i].Date].push(data[i]);
      }
    }
    var totalTypeData = this.createTypeColumns(data);
    var types = {};
    for (var i = 0; i < totalTypeData.length; i++) {
      var key = totalTypeData[i][0];
      if (types[key] === undefined) {
        types[key] = [];
      }
    }
    for (key in obj) {
      x.push(moment(key).format("YYYY-MM-DD"));
      var dayTypeData = this.createTypeColumns(obj[key]);
      for (typeKey in types) {
        var count = 0;
        for (var j = 0; j < dayTypeData.length; j++) {
          var type = dayTypeData[j][0];
          if (typeKey === type) {
            count = dayTypeData[j][1];
          }
        }
        types[typeKey].push(count);
      }
    }

    var result = [x];
    for (key in types) {
      types[key].unshift(key.slice(0, 9));
      result.push(types[key]);
    }
    return result;
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
