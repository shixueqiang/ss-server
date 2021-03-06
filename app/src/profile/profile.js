import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import ProfileList from './profileList.js'
import ProfileEdit from './profileEdit.js'
import ProfileInsert from './profileInsert.js'
import { Router, Route, Link, BrowserRouter } from 'react-router-dom'
import '../index.css';
import App from '../App';
class ProfileIndex extends Component {
    render() {
        return (
        <BrowserRouter>
            <div>
                {/* <Route path="/" component={ProfileList}/> */}
                <Route path="/toProfiles" component={ProfileList}/>
                <Route path="/toProfileEdit" component={ProfileEdit} />
                <Route path="/toProfileInsert" component={ProfileInsert} />
              </div>
        </BrowserRouter>
        );
    }
}

ReactDOM.render(<ProfileIndex />, document.getElementById('root'));