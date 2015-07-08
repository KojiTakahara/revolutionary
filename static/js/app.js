"use strict";

var app = angular.module('app', [
  "720kb.datepicker",
  "chartService",
  "c3Service"
]);

app.controller('indexCtrl', ['$scope', '$http', '$filter', '$sce', '$window', "chartService", "c3Service", function($scope, $http, $filter, $sce, $window, chartService, c3Service) {

  $scope.data = [];
  $scope.startDate = moment().add(-7, "days").format('YYYY-MM-DD');
  $scope.endDate = moment().format('YYYY-MM-DD');

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
      $scope.viewRace();
    });
  };

  $scope.viewType = function() {
    var columns = chartService.createTypeColumns($scope.data);
    $scope.chart = c3Service.drowDonutChart("#typeChart", columns, "deck type");
  };

  $scope.viewRace = function() {
    var columns = chartService.createRaceColumns($scope.data);
    $scope.chart = c3Service.drowDonutChart("#typeRace", columns, "deck race");
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
