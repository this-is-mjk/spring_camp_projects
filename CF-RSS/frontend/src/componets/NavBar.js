// Navbar.js
import React from 'react';
import './NavBar.css';
import { useNavigate } from 'react-router-dom';

const Navbar = ({ isLoggedIn }) => {
  isLoggedIn = false
  const navigate = useNavigate();

  const signin = () => {
    navigate(`/signin`);
  };

  const signup = () => {
    navigate(`/signup`);
  };

  const onRefresh = () => {
    window.location.reload();
  };

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
    </div>
  );
};

export default Navbar;
