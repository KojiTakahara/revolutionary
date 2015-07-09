var service = angular.module('c3Service', []);
service.factory('c3Service', ['$http', '$q', function($http, $q) {
  var service = {};

  /**
   * ドーナツグラフを描く
   * @param bindTo 描画対象のID
   * @param columns 内容
   * @param title タイトル
   */
  service.drowDonutChart = function(bindTo, columns, title) {
    return c3.generate({
      bindto: bindTo,
      data: { columns: columns, type : "pie" },
      donut: { title: title }
    });
  };

  service.drowTimeseriesChart = function(bindTo, columns) {
    return c3.generate({
      bindto: bindTo,
      data: {
        x: 'x',
        columns: columns
      },
      axis: {
        x: {
          type: 'timeseries',
          tick: {
            format: '%Y-%m-%d'
          }
        }
      }
    });
  };

  return service;
}]);
