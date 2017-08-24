import React, {Component} from 'react';
import ReactDOM from 'react-dom';
import ProfileList from './profileList.js'
import ProfileEdit from './profileEdit.js'
import { Router, Route, Link, BrowserRouter } from 'react-router-dom'
import '../index.css';
import App from '../App';
const ProfileIndex = () => (
    <BrowserRouter>
    <div>
        {/* <Route path="/" render={() => <ProfileList />}/> */}
        <Route path="/profiles" render={() => <ProfileList />}/>
        <Route path="/profileEdit" render={() => <ProfileEdit />} />
      </div>
    </BrowserRouter>
);

ReactDOM.render(<ProfileIndex />, document.getElementById('root'));