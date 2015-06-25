"use strict";

var app = angular.module('app', [
]);

app.controller('indexCtrl', ['$scope', '$http', '$sce', '$window', function($scope, $http, $sce, $window) {

  $scope.submit = function() {
    $http({
      method: "GET",
      url: "/api/tournament_history",
      param: {
        // startDate: $scope.startDate,
        // endDate: $scope.endDate,
        // count: $scope.count
      }
    }).success(function(data) {
      console.log(data);
    }).error(function(data, status, headers, config) {
    });

    $scope.chart = c3.generate({
      bindto: '#chart',
      data: {
        columns: [
          ['data1', 30, 200, 100, 400, 150, 250],
          ['data2', 50, 20, 10, 40, 15, 25]
        ]
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