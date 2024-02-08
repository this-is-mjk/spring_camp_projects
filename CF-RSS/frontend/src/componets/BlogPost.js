import React from 'react';
import {useGlobalVar} from './logindetails'
import { useNavigate } from 'react-router-dom';

const BlogPost = ({ title, userName, postingTime, id}) => {
  const navigate = useNavigate();
  const controller = new AbortController();
  const { emailid, password, isLoggedIn, SubscribedId } = useGlobalVar();
  const handleCommentsClick = () => {
    // Implement logic to handle comments click
    console.log('View comments clicked!');
    console.log(id);
  };

  const handleSubscribeClick = (id) => {
    console.log('Subscribe clicked!');
    if (!isLoggedIn) {
      alert("please Login!")
      navigate('/signin')
      return
    }
    controller.abort();
    fetch('http://localhost:8080/user/blogs/subscribe', {
      method: 'POST',
      body: JSON.stringify({
        email: emailid,
        password: password,
        blogid: id
      }),
      headers: {
        'Content-Type': 'application/json'
      }
    }).then((res) => {
        if (res.status === 250) {
            alert("User not found/incorrect user - login again!")
            return
        }
    }).catch((error) => {
      console.log(error)
      alert("Error occured, please try again.");
    });
  };
  return (
    <div className="blog-post">
      <h2>{title}</h2>
      <p>By {userName} on {postingTime}</p>
      <div className="button-div">
        <button onClick={handleCommentsClick}>View Comments</button>
        { isLoggedIn &&
          <button onClick={() => { 
            handleSubscribeClick(id);
            }}>Subscribe</button>
        }
        
      </div>
    </div>
  );
};

export default BlogPost;

// SubscribedId.indexOf(id) !== -1