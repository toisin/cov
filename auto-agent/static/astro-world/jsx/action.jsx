/** @jsx React.DOM */

// npm install -g react-tools
// jsx -w -x jsx public/js public/js


var Action = React.createClass({

  getInitialState: function() {
    return {mode: 0};
  },

  render: function() {
    var state = this.state;
    var user = this.props.user;
    var app = this.props.app;
    var prompt = user.getPrompt();
    // debugger;
    switch (user.getCurrentPhaseId()) {
      case PHASE_COV:
        return  <div><CovAction user={user} prompt={prompt} onComplete={app.changeState} app={app}/></div>;
      // case "chart":
      //   return <div></div>
      // case "prediction":
      //   return <div></div>

    //     break;
    }
  }
});


