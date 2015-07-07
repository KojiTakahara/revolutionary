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
      data: { columns: columns, type : "donut" },
      donut: { title: title }
    });
  };

  return service;
}]);
