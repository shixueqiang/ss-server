import React, {Component} from 'react';
import TextField from 'material-ui/TextField';
import RaisedButton from 'material-ui/RaisedButton';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import $ from 'jquery';
import Dialog from 'material-ui/Dialog';
import FlatButton from 'material-ui/FlatButton';

const style = {
    marginLeft: 60,
    marginTop: 12,
    marginBottom: 12,
};

class ProfileInsert extends Component {
    state = {
        open: false,
        result: "",
    };

    handleOpen = () => {
        this.setState({open: true,});
    };
    
    handleClose = () => {
        this.setState({open: false,});
        this.props.history.push({pathname:'/toProfiles'});
    };

    handleClick = () => {
        const _this = this;
        $.ajax({
            url:"/profileInsert",
            data:$("#myform").serialize(),
            type:"POST",
            success: function(data) {
                console.log(data);
                _this.setState({open:true,result: data,});
            }
        });
    }

    render() {
        const actions = [
            <FlatButton
              label="OK"
              primary={true}
              onClick={this.handleClose}
            />,
          ];
        return(
        <div>
            <MuiThemeProvider>
                <div>
                    <form id="myform" action="/profileInsert" method="post">
                        <TextField className="profileTextField" floatingLabelText="Name" name="Name"/><br />
                        <TextField className="profileTextField" floatingLabelText="Host" name="Host"/><br />
                        <TextField className="profileTextField" floatingLabelText="LocalPort" defaultValue="1080" name="LocalPort"/><br />
                        <TextField className="profileTextField" floatingLabelText="RemotePort" defaultValue="8388" name="RemotePort"/><br />
                        <TextField className="profileTextField" floatingLabelText="Password" type="password" name="Password"/><br />
                        <TextField className="profileTextField" floatingLabelText="Protocol" defaultValue="origin" name="Protocol"/><br />
                        <TextField className="profileTextField" floatingLabelText="ProtocolParam" name="ProtocolParam"/><br />
                        <TextField className="profileTextField" floatingLabelText="Obfs" defaultValue="plain" name="Obfs"/><br />
                        <TextField className="profileTextField" floatingLabelText="ObfsParam" name="ObfsParam"/><br />
                        <TextField className="profileTextField" floatingLabelText="Method" defaultValue="aes-256-cfb" name="Method"/><br />
                        <TextField className="profileTextField" floatingLabelText="Route" defaultValue="all" name="Route"/><br />
                        <TextField className="profileTextField" floatingLabelText="RemoteDNS" defaultValue="8.8.8.8" name="RemoteDNS"/><br />
                        <TextField className="profileTextField" floatingLabelText="VpnType" defaultValue="1" name="VpnType"/><br />
                        <TextField className="profileTextField" floatingLabelText="Ikev2Type" defaultValue="1" name="Ikev2Type"/><br />
                        <RaisedButton label="submit" primary={true} style={style}  onClick={this.handleClick}/>
                    </form>
                    <Dialog
                        actions={actions}
                        modal={false}
                        open={this.state.open}
                        onRequestClose={this.handleClose}
                        >
                        {this.state.result}
                    </Dialog>
                </div>
            </MuiThemeProvider>
        </div>
        );
    };
}
export default ProfileInsert;