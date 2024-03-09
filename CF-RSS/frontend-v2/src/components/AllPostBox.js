import React from "react";
import "./AllPostBox.css";
import { useState, useEffect } from "react";
import OnePost from "./OnePost"
import axios from 'axios'
import { useNavigate} from 'react-router-dom';
import _ from 'lodash';

export default function AllPostBox (){
    const [data, setData] = useState(null);
    const [groupedData, setGroupedData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [locallySaved, setlocallSaved] = useState(false);
    const [subscriptions, setSubscriptions] = useState(null);
    const navigate = useNavigate();
    const Url = 'http://localhost:8080/all-recent-actions';

    useEffect(() => {
        axios.get(Url, {withCredentials: true}) 
            .then(response => {
                console.log('data fetch successful!');
                setData(response.data.posts);
                setSubscriptions(response.data.subscriptions);
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
                }
            });
            setLoading(false)
    }, []);
    useEffect(() => {
        if(data){
            sessionStorage.setItem('groupedCommentData', JSON.stringify(_.groupBy(data, (recentAction) => recentAction.blogEntry.id)));
            setlocallSaved(true);
        }
    }, [data]);
    useEffect(() => {
        sessionStorage.setItem("subscriptions", JSON.stringify(subscriptions));
    }, [subscriptions])
    useEffect(()=>{
        const items = JSON.parse(sessionStorage.getItem('groupedCommentData'));
        if (items) {
            // // Convert the object into an array of [key, value] pairs
            // const arrayData = Object.entries(items).map(([key, value]) => ({ key, value }));
            setGroupedData(items);
        }
    }, [locallySaved])

    return (
        <div className="all-post-box">
        <div className="bg-image-all-post"></div>
            {loading && <h1 className="heading">loading...</h1>}
            {!loading && <h1 className="heading">Recent Posts</h1>}
            <div className="top-layer">
            {subscriptions && groupedData && Object.entries(groupedData).map((post) => {
                // post is an array [postId, actions]
                // in which actions is also a array of actions
                // console.log(post[0]) -- postId
                // console.log(post[1]) -- array of actions
                // post[1].map((action) => {console.log(action.blogEntry.title)}) -- will give all title of all actions but we only need once
                // console.log(post[1][0].blogEntry.title.slice(3,-4))
                return (<OnePost 
                    key={post[0]} 
                    title={post[1][0].blogEntry.title.slice(3,-4)}
                    postId={post[0]}
                    userSubscriptions = {JSON.stringify(subscriptions)}
                    />)
            })}
            </div>
        </div>
    );
};
// {data.blogentry.title}