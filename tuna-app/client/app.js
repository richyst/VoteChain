// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	// $scope.queryAllTuna = function(){

	// 	appFactory.queryAllTuna(function(data){
	// 		var array = [];
	// 		for (var i = 0; i < data.length; i++){
	// 			parseInt(data[i].Key);
	// 			data[i].Record.Key = parseInt(data[i].Key);
	// 			array.push(data[i].Record);
	// 		}
	// 		array.sort(function(a, b) {
	// 		    return parseFloat(a.Key) - parseFloat(b.Key);
	// 		});
	// 		$scope.all_tuna = array;
	// 	});
	// }

	// $scope.queryTuna = function(){

	// 	var id = $scope.tuna_id;

	// 	appFactory.queryTuna(id, function(data){
	// 		$scope.query_tuna = data;

	// 		if ($scope.query_tuna == "Could not locate tuna"){
	// 			console.log()
	// 			$("#error_query").show();
	// 		} else{
	// 			$("#error_query").hide();
	// 		}
	// 	});
	// }

	// $scope.recordTuna = function(){

	// 	appFactory.recordTuna($scope.tuna, function(data){
	// 		$scope.create_tuna = data;
	// 		$("#success_create").show();
	// 	});
	// }

	// ----------------*******************---------------------------------

	$scope.queryAllVotes = function(){

		appFactory.queryAllVotes(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_votes = array;
		});
	}

	$scope.queryVote = function(){

		var id = $scope.vote_id;

		appFactory.queryVote(id, function(data){
			$scope.query_vote = data;

			if ($scope.query_vote == "Could not locate vote"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.queryVotesByEleccion = function(){

		var eleccion = $scope.vote_eleccion;

		appFactory.queryVotesByEleccion(eleccion, function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});

			$scope.query_vote_eleccion = array;

			if ($scope.query_vote_eleccion == "Could not locate votes"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordVote = function(){

		appFactory.recordVote($scope.vote, function(data){
			$scope.create_vote = data;
			$("#success_create").show();
		});
	}

	// --------------------------------***************************************

	// $scope.changeHolder = function(){

	// 	appFactory.changeHolder($scope.owner, function(data){
	// 		$scope.change_holder = data;
	// 		if ($scope.change_holder == "Error: no tuna catch found"){
	// 			$("#error_holder").show();
	// 			$("#success_holder").hide();
	// 		} else{
	// 			$("#success_holder").show();
	// 			$("#error_holder").hide();
	// 		}
	// 	});
	// }

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    // factory.queryAllTuna = function(callback){

    // 	$http.get('/get_all_tuna/').success(function(output){
	// 		callback(output)
	// 	});
	// }

	// factory.queryTuna = function(id, callback){
    // 	$http.get('/get_tuna/'+id).success(function(output){
	// 		callback(output)
	// 	});
	// }

	// factory.recordTuna = function(data, callback){

	// 	data.location = data.longitude + ", "+ data.latitude;

	// 	var tuna = data.id + "-" + data.location + "-" + data.timestamp + "-" + data.owner + "-" + data.vessel;

    // 	$http.get('/add_tuna/'+tuna).success(function(output){
	// 		callback(output)
	// 	});
	// }

	//-------------------------------------------------------------

	factory.queryAllVotes = function(callback){

    	$http.get('/get_all_votes/').success(function(output){
			callback(output)
		});
	}

	factory.queryVote = function(id, callback){
    	$http.get('/get_vote/' + id).success(function(output){
			callback(output)
		});
	}

	factory.queryVotesByEleccion = function(eleccion, callback){
    	$http.get('/get_by_eleccion/' + eleccion).success(function(output){
			callback(output)
		});
	}

	factory.recordVote = function(data, callback){

		var vote = data.id + "-" + data.codigo + "-" + data.eleccion + "-" + data.entidad + "-" + data.distrito
			+ "-" + data.municipio + "-" + data.seccion + "-" + data.localidad;

    	$http.get('/add_vote/' + vote).success(function(output){
			callback(output)
		});
	}

	//--------------------------------------------------------

	// factory.changeHolder = function(data, callback){

	// 	var owner = data.id + "-" + data.name;

    // 	$http.get('/change_holder/' + owner).success(function(output){
	// 		callback(output)
	// 	});
	// }

	return factory;
});


