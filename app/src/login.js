import React from 'react';
import ReactDOM from 'react-dom';
import './App.css';
import $ from 'jquery';
import TextField from 'material-ui/TextField';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';

class NormalLoginForm extends React.Component {

  handleClick = () => {
    const _this = this;
    $.ajax({
        url:"/login",
        data:$("#myform").serialize(),
        type:"POST",
        success: function(data) {
            console.log(data);
            if(data === "success")
              window.location.href="/toProfiles";
        }
    });
  }

  render() {
    return (
      <MuiThemeProvider>
        <form id="myform" action="/login" method="post">
          <TextField floatingLabelText="Account" name="account"/>
          <TextField floatingLabelText="Password" name="password" type="password"/>
          <button type="button" className="button buttonBlue" onClick={this.handleClick}>LOGIN</button>
        </form>
      </MuiThemeProvider>
    )
  }
}

ReactDOM.render(<NormalLoginForm />, document.getElementById('root'));