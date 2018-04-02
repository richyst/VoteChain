//SPDX-License-Identifier: Apache-2.0

var tuna = require('./controller.js');
var cors = require('cors');

module.exports = function(app){
  app.use(cors());
  app.get('/get_tuna/:id', function(req, res){
    tuna.get_tuna(req, res);
  });
  app.get('/add_tuna/:tuna', function(req, res){
    tuna.add_tuna(req, res);
  });
  app.get('/get_all_tuna', function(req, res){
    tuna.get_all_tuna(req, res);
  });
  app.get('/change_holder/:owner', function(req, res){
    tuna.change_holder(req, res);
  });

  // ---------------------------------------

  app.get('/get_vote/:id', function(req, res){
    tuna.get_vote(req, res);
  });

  app.get('/get_by_domain/:search', function(req, res){
    tuna.get_by_domain(req, res);
  });

  app.get('/add_vote/:vote', function(req, res){
    tuna.add_vote(req, res);
  });

  app.get('/get_all_votes', function(req, res){
    tuna.get_all_votes(req, res);
  });

}
