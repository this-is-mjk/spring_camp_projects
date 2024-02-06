// Navbar.js
import React, { Fragment, useState, useContext } from 'react'
import './NavBar.css';
import { useNavigate } from 'react-router-dom';
import {useGlobalVar} from './logindetails'

const Navbar = () => {
  const { emailid, password, isLoggedIn, updateEmail, updatePassword, updateIsLoggedIn } = useGlobalVar();
  const navigate = useNavigate();
  // isLoggedIn = false
    const signin = () => {
      navigate(`/signin`);
    };

    const signup = () => {
      navigate(`/signup`);
    };

    const onRefresh = () => {
      window.location.reload();
    }
    const logout = () => {
      updateEmail('')
      updatePassword('')
      updateIsLoggedIn(false)
    }
      return (
      <div className="navbar">
        <button className="navbar-button" onClick={onRefresh}>
          Refresh
        </button>
        {!isLoggedIn && (
          <div className="right">
            <button className="navbar-button" onClick={signin}>
              Login
            </button>
            <button className="navbar-button" onClick={signup}>
              Sign Up
            </button>
          </div>
        )}
        {isLoggedIn && (
          <div className="right">
            <button className="navbar-button" onClick={logout}>
              Logout
            </button>
          </div>
        )}
      </div>
  );
};

export default Navbar;
