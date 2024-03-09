import React, { useState, useEffect } from "react";
import "./ProfilePage.css";
import { IoMdContact } from "react-icons/io";
import OnePost from "./OnePost";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import Cookies from 'js-cookie';

export default function ProfilePage () {
    const [groupedData, setGroupedData] = useState(null); 
    const [subscriptions, setSubscriptions] = useState(null);
    const [loading, setLoading] = useState(true);
    const [islocallySaved, setlocallySaved] = useState(false);
    const navigate = useNavigate();
    const Url = 'http://localhost:8080/all-recent-actions'
    useEffect(()=>{
        const items = JSON.parse(sessionStorage.getItem('groupedCommentData'));
        if (items) {
            // // Convert the object into an array of [key, value] pairs
            // const arrayData = Object.entries(items).map(([key, value]) => ({ key, value }));
            setGroupedData(items);
        }
    }, [])
    useEffect(() =>{
        sessionStorage.setItem("subscriptions", JSON.stringify(subscriptions));
        setlocallySaved(true)
    }, [subscriptions])
    useEffect(() => {
        axios.get(Url, {withCredentials: true}) 
            .then(response => {
                console.log('data fetch successful!');
                setSubscriptions(response.data.subscriptions);
                setLoading(false);
            })
            .catch(error => {
                if (error.response && error.response.status === 401) {
                    // The server responded with a 400 status code
                    // alert("Please Sign In again!");
                    console.log('unauthorised');
                    navigate("/login");
                }
                else {
                    alert('Error fetching posts, please try again.');
                    console.log(error);
                    setLoading(false);
                }
            });
    }, []);
    // return(
    //     <div className="bgImageProfilePage">
    //         <div className="left-size-contact-card">
    //             <IoMdContact className="user-image"/>
    //             <div className="user-details">
    //                 <p>Username: this_is_mjk</p>
    //                 <p>Email: mjlivetowork@gmail.com</p>
    //             </div>
    //         </div>
    //         <div className="right-size-subscriptions">
    //             <h1 className="profile-subscribe-heading" style={{color: "white", fontStyle: "italic"}}>Subscribed</h1>
    //             {loading ? (
    //                 <p>Loading...</p> // Show a loading indicator
    //             ) : (
    //                 <div className="posts-sub">
    //                 {subscriptions && groupedData && Object.entries(groupedData).map((post) => {
    //                     if (subscriptions.includes(post[0])) {
    //                     return (
    //                         <OnePost
    //                         key={post[0]}
    //                         title={post[1][0].blogEntry.title.slice(3, -4)}
    //                         postId={post[0]}
    //                         userSubscriptions={JSON.stringify(sessionStorage.getItem("subscriptions") ?? [])}
    //                         />
    //                     );
    //                     }
    //                 })}
    //                 </div>)}
    //         </div>
    //     </div>
    // )
    return (
        <div className="bgImageProfilePage">
          <div className="left-size-contact-card">
            <IoMdContact className="user-image" />
            <div className="user-details">
              <p>Username: {Cookies.get('Username')}</p>
              <p>Email: {Cookies.get('Email')}</p>
            </div>
          </div>
          <div className="right-size-subscriptions">
            <h1 className="profile-subscribe-heading" style={{ color: "white", fontStyle: "italic" }}>Subscribed</h1>
            {loading ? (
              <p>Loading...</p>
            ) : (
              <div className="posts-sub">
                {islocallySaved && groupedData && Object.entries(groupedData).map((post) => {
                    if(sessionStorage.getItem("subscriptions").includes(post[0])){
                        return (
                            <OnePost
                              key={post[0]}
                              title={post[1][0].blogEntry.title.slice(3, -4)}
                              postId={post[0]}
                              userSubscriptions={sessionStorage.getItem("subscriptions")}
                            />
                          );
                    }
                })}
              </div>
            )}
          </div>
        </div>
      );
}