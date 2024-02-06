import React from 'react';
import {useEffect, useState} from 'react';
import NavBar from './NavBar';
import BlogPost from './BlogPost'
import './BlogPost.css';
const addToBlogBox = (responseData) => {
}
const PostSection = () => {
    const [blog, setBlog] = useState();
    const [isLoading, setIsLoading] = useState(false);
    const [isError, setIsError] = useState(false);
    useEffect(() => {
        const fetchPost = () => {
            setIsLoading(true);
            fetch('http://localhost:8080/activity/recent-actions').then((response) => {
                if (response.ok) {
                    return response.json();
                } else {
                    throw 'Error getting users list'
                }
            }).then((j) => {
                setBlog(j);
            }).catch((error) => {
                setIsError(true)
            })
            setIsLoading(false);
        }
        fetchPost();
        
    }, [process.env.REACT_APP_LoggedIn]);
    return (
        <div id="blogPage">
            <NavBar />
            <h1>Community</h1>
            {blog && blog.map((ele, index) => 
                <BlogPost 
                key={index}
                title={ele.blogEntry.title.slice(3,-4)}
                userName={ele.blogEntry.authorHandle}
                postingTime={ele.timeSeconds}
              />
            )}
        </div>
      );
};

export default PostSection;
