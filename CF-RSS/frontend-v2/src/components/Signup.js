import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import axios from 'axios';
import './LoginSignup.css';
import bcrypt from 'bcryptjs'

export default function Signup () {
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [isValidPassword, setValidPassword] = useState(true);
    const [isValidUsername, setValidUsername] = useState(true);
    const [isValidEmail, setValidEmail] = useState(true);
    localStorage.clear()
    sessionStorage.clear()
    const pastDate = new Date(1970, 0, 1).toUTCString();
    document.cookie = "Authorization=; expires=" + pastDate + "; path=/;";
    document.cookie = "Username=; expires=" + pastDate + "; path=/;";
    document.cookie = "Email=; expires=" + pastDate + "; path=/;";

    const handleChangeUsername = (event) => {
        const input = event.target.value;
        setUsername(input);
        setValidUsername(input.trim() !== ''); // Basic validation: Username should not be empty
    };
    const handleChangePassword = (event) => {
        const input = event.target.value;
        setPassword(input);
        // Basic validation: Password should have a minimum length of 6 characters
        setValidPassword(input.length >= 6);
    };
    const handleChangeEmail = (event) => {
        const input = event.target.value;
        setEmail(input);
        setValidEmail(emailPattern.test(input));
    }
    const handleSignup = () => {
        // server
        const Url = 'http://localhost:8080/signup';
        
        if (password.length < 6 || username.length < 1 || email.length < 1 || !isValidEmail) {
            alert("Please enter valid Username, Password, email")
            return
        }

        // Data to be sent in the POST request
        const postData = {
            // hash the password
            "password": bcrypt.hashSync(password, '$2a$10$CwTycUXWue0Thq9StjUM0u'),
            "username": username,
            "email": email,
            "subscriptions": [],
        };

        // Sending a POST request using axios
        // Axios automatically transforms request and response data to JSON. 
        // It simplifies working with JSON data as we don't need to explicitly parse responses.
        axios.post(Url, postData)
            .then(response => {
                console.log('Login successful!', response.data);
                // Handle successful signup
                navigate("/login");
                setEmail("");
                setPassword("");
                setUsername("");
                alert("Thank you, now please login to your new profile!")
            })
            .catch(error => {
                if (error.response && error.response.status === 400) {
                    // The server responded with a 400 status code
                    alert(error.response.data);
                console.log('SignUp failed!');
                }
                else{
                    console.log(error)
                }
            });
    };  
    return (
        <div className='bg-image'>
            <div className='login-container'>
                <div id='email' className="form-group">
                    <label htmlFor="email">Email: </label>
                    <input
                        type="email"
                        id="username"
                        value={email}
                        onChange={handleChangeEmail}
                    />
                    {!isValidEmail && <p style={{ color: 'red' ,margin: '0'}}><small>Not valid email</small></p>}
                </div>
                <div id='username' className="form-group">
                    <label htmlFor="username">Username: </label>
                    <input
                        type="text"
                        id="username"
                        value={username}
                        onChange={handleChangeUsername}
                    />
                    {!isValidUsername && <p style={{ color: 'red' ,margin: '0'}}><small>Username cannot be empty</small></p>}
                </div>
                <div id='password' className="form-group">
                    <label htmlFor="password">Password: </label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={handleChangePassword}
                    />
                    {!isValidPassword && <p style={{ color: 'red', margin: '0'}}><small>Password should be at least 6 characters long</small></p>}
                </div>
                <button id='submitBtn' onClick={handleSignup}>SignUp</button>
                <Link to = "/"><p style={{display : "inline", paddingLeft : "20px", color : "white"}}>Already have a account?</p></Link>
            </div>
        </div>
        );
}

