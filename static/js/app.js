"use strict";

var app = angular.module('app', [
  "720kb.datepicker"
]);

app.controller('indexCtrl', ['$scope', '$http', '$filter', '$sce', '$window', function($scope, $http, $filter, $sce, $window) {

  $scope.data = [];

  $scope.startDate = $filter('date')(new Date(), "yyyy-MM-dd");
  $scope.endDate = $filter('date')(new Date(), "yyyy-MM-dd");

  $scope.submit = function() {
    $http({
      method: "GET",
      url: "/api/tournament_history",
      params: {
        startDate: $scope.startDate,
        endDate: $scope.endDate,
        count: $scope.count
      }
    }).success(function(data) {
      $scope.data = data;
      $scope.viewType();
    });
  };

  $scope.viewType = function() {
    var obj = {};
    for (var i = 0; i < $scope.data.length; i++) {
      if ($scope.data[i].Type === "") {
        continue;
      }
      if (obj[$scope.data[i].Type] === undefined) {
        obj[$scope.data[i].Type] = 1;
      } else {
        obj[$scope.data[i].Type]++;
      }
    }
    var columns = [];
    for (var key in obj) {
      columns.push([key, obj[key]]);
    }
    typeDonut(columns, "deck type");
  };

  $scope.viewRace = function() {
    var obj = {};
    for (var i = 0; i < $scope.data.length; i++) {
      if ($scope.data[i].Race === "") {
        continue;
      }
      if (obj[$scope.data[i].Race] === undefined) {
        obj[$scope.data[i].Race] = 1;
      } else {
        obj[$scope.data[i].Race]++;
      }
    }
    var columns = [];
    for (var key in obj) {
      columns.push([key, obj[key]]);
    }
    typeDonut(columns, "deck race");
  };

  $scope.viewColor = function() {
    var light = ["light"],
        water = ["water"],
        dark = ["dark"],
        fire = ["fire"],
        nature = ["nature"],
        zero = ["zero"];
    for (var i = 0; i < $scope.data.length; i++) {
      if ($scope.data[i].Light) {
        light.push(1);
      }
      if ($scope.data[i].Water) {
        water.push(1);
      }
      if ($scope.data[i].Dark) {
        dark.push(1);
      }
      if ($scope.data[i].Fire) {
        fire.push(1);
      }
      if ($scope.data[i].Nature) {
        nature.push(1);
      }
      if ($scope.data[i].Zero) {
        zero.push(1);
      }
    }
    typeDonut([light, water, dark, fire, nature, zero], "deck color");
  };

  var typeDonut = function(columns, title) {
    $scope.chart = c3.generate({
      bindto: '#typeChart',
      data: {
        columns: columns,
        type : "donut"
      },
      donut: {
        title: title
      }
    });
  };

}]);

app.config(['$httpProvider', '$locationProvider', function($httpProvider, $locationProvider) {
  $httpProvider.defaults.headers.common = {'X-Requested-With': 'XMLHttpRequest'};
  $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded; charset=utf-8';
  $httpProvider.defaults.transformRequest = function(data) {
    if (data === undefined) {
      return data;
    }
    return $.param(data);
  }
  $locationProvider.html5Mode({
    enabled: true,
    requireBase: false
  });
}]);
