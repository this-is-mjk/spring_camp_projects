// Import React and CSS for styling
import React from 'react';
import NavBar from './NavBar';
import {useGlobalVar} from './logindetails'
import {useEffect, useState} from 'react';
import BlogPost from './BlogPost'
// Functional component for the empty page
const Subscribed = () => {
    const { emailid, password, isLoggedIn} = useGlobalVar();
    const [blog, setBlog] = useState();
    const [isLoading, setIsLoading] = useState(false);
    const [isError, setIsError] = useState(false);
    const controller = new AbortController();
        // fetch('http://localhost:8080/user/activity/recent-actions', {
        //     method: 'POST',
        //     body: JSON.stringify({
        //         email: emailid,
        //         password: password,
        //       }),
        //       headers: {
        //         'Content-Type': 'application/json'
        //       }
        // }).then((resp) => {

        // })
        useEffect(() => {
            if (!isLoggedIn){
                return (
                    <h1>Please Login First</h1>
                )
            }
            controller.abort();
            const fetchPost = () => {
                setIsLoading(true);
                fetch('http://localhost:8080/user/activity/recent-actions', {
                    method: 'POST',
                    body: JSON.stringify({
                        email: emailid,
                        password: password,
                      }),
                      headers: {
                        'Content-Type': 'application/json'
                      }
                }).then((response) => {
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
            
        }, []);
    return (
        <div>
            <NavBar />
            <h1 >Subscriptions</h1>
            {blog && blog.map((ele, index) => 
                <BlogPost 
                key={index}
                title={ele.blogEntry.title.slice(3,-4)}
                userName={ele.blogEntry.authorHandle}
                postingTime={ele.timeSeconds}
                id={ele.blogEntry.id}
              />
            )}
        </div>
    );
};
export default Subscribed;
