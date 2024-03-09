import React, { useEffect, useState} from 'react';
import { useNavigate, Link } from 'react-router-dom';
import axios from 'axios';
import './LoginSignup.css';
import bcrypt from 'bcryptjs'

export default function Login () {
    const navigate = useNavigate();
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isValidPassword, setValidPassword] = useState(true);
    const [isValidUsername, setValidUsername] = useState(true);
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
    const handleLogin = () => {
        // server
        const Url = 'http://localhost:8080/login';
        
        if (password.length < 6 || username.length < 1) {
            alert("Please enter Username and Password")
            return
        }
        // Data to be sent in the POST request
        const postData = {
            // hash the password
            "password": bcrypt.hashSync(password, '$2a$10$CwTycUXWue0Thq9StjUM0u'),
            "username": username
        };

        // Sending a POST request using axios
        // Axios automatically transforms request and response data to JSON. 
        // It simplifies working with JSON data as we don't need to explicitly parse responses.

        // request part
        axios.post(Url, postData, {withCredentials: true})
            .then(response => {
                console.log('Login successful!', response.data);
                // Handle successful login
                navigate("/");
            })
            .catch(error => {
                if (error.response && error.response.status === 400) {
                    // The server responded with a 400 status code
                    alert(error.response.data);
                    console.log('Login failed!');
                }
                else{
                    console.log(error)
                    alert("Error occured please try again later.")
                }
            });
    };
    return (
        <div className='bg-image'>
            <div className='login-container'>
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
                <button id='submitBtn' onClick={handleLogin}>Login</button>
                <Link to = "/signup"><p style={{display : "inline", paddingLeft : "20px", color : "white"}}>Don't have a account?</p></Link>
            </div>
        </div>
        );
}