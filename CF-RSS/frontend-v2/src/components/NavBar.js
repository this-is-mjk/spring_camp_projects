// Navbar.js

import React, { useState } from "react";
import { Link } from "react-router-dom";
import "./NavBar.css"; // Import the corresponding CSS file

const Navbar = () => {
  const [menuOpen, setMenuOpen] = useState(false);

  const toggleMenu = () => {
    setMenuOpen(!menuOpen);
  };

  return (
    <nav className={`navbar ${menuOpen ? "mobile" : ""}`}>
      <div className="navbar-content">
        <div className={`navbar-links ${menuOpen ? "open" : ""}`}>
          <Link to="/">All Posts</Link>
          <Link to="/myprofile">My Profile</Link>
          <Link to="/login">Log Out</Link>
        </div>
        <div className="mobile-menu" onClick={toggleMenu}>
          <div className={`bar ${menuOpen ? "open" : ""}`}></div>
          <div className={`bar ${menuOpen ? "open" : ""}`}></div>
          <div className={`bar ${menuOpen ? "open" : ""}`}></div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
