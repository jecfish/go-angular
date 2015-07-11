var demoApp = angular.module('demoApp',['ngRoute']);

demoApp.config(['$routeProvider', function($routeProvider){
  $routeProvider
  .when('/',{
    templateUrl:'home.html'
  })
  .otherwise({
    redirectTo: '/'
  });
}]);

demoApp.controller('demoCtrl', function(demoFactory){
  var vm = this;

  vm.name = 'who';

  vm.list = demoFactory.friends;

  vm.create = function(name) {
    demoFactory.createFriend(name);
  };

  vm.remove = function(name){
    demoFactory.deleteFriend(name);
  }

  demoFactory.getFriends().then(function(){
    vm.list = demoFactory.friends;
  });

});

demoApp.factory('demoFactory', function($http){
  var model = {};
  model.friends = [];

  model.createFriend = function(friend){
    $http.post('http://localhost:8000/friend/' + friend).then(function(){
      model.friends.push(friend);
    });
  };

  model.getFriends = function(){
    return $http.get('http://localhost:8000/friend/').success(
      function(data) {
        model.friends = data;
      });
    };

    model.deleteFriend = function(friend){
      $http.delete('http://localhost:8000/friend/' + friend).then(function(){
        var index = model.friends.indexOf(friend);
        model.friends.splice(index,1);
      });
    };

    return model;
  });
