import React from 'react';

const handleCommentsClick = () => {
    // Implement logic to handle comments click
    console.log('View comments clicked!');
};

const handleSubscribeClick = () => {
    // Implement logic to handle subscribe click
    console.log('Subscribe clicked!');
};


const BlogPost = ({ title, userName, postingTime}) => {
  return (
    <div className="blog-post">
      <h2>{title}</h2>
      <p>By {userName} on {postingTime}</p>
      <div class="button-div">
        <button onClick={handleCommentsClick}>View Comments</button>
        <button onClick={handleSubscribeClick}>Subscribe</button>
      </div>
    </div>
  );
};

export default BlogPost;
