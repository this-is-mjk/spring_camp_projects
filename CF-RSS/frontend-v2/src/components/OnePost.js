import React from "react";
import "./AllPostBox.css";
// import AllCommentsPage from "./AllCommentsPage"
import { useNavigate } from "react-router-dom";
import axios from "axios";
export default function OnePost (props) {
    const Url = "http://localhost:8080/subscribe-request"
    const navigate = useNavigate();
    const openCommentPage = () => {
        // alert("all comments clicked");
        return (
                navigate("/AllCommentsPage", {state: {postId: props.postId}})
        )
    }
    const Subscribe = () => {
        const postdata = {
            "postid": parseInt(props.postId),
            "subscribe": true
        }
        axios.post(Url, postdata, {withCredentials: true}) 
            .then(() => {
                console.log('subscribe successful!');
                document.getElementById(`"subsBtn-${props.postId}"`).style.display = "none";
                document.getElementById(`"unsubsBtn-${props.postId}"`).style.display = "block";
            }).catch(error => {
                if (error.response && error.response.status === 401) {
                    // The server responded with a 400 status code
                    alert('please login again');
                    navigate("/login");
                }
                else{
                    console.log(error)
                    alert("Error occured please try again later.")
                }
            });
    }
    const Unsubscribe = () => {
        const postdata = {
            "postid": parseInt(props.postId),
            "subscribe": false
        }
        axios.post(Url, postdata, {withCredentials: true}) 
            .then(() => {
                console.log('unsubscribe successful!');
                document.getElementById(`"subsBtn-${props.postId}"`).style.display = "block";
                document.getElementById(`"unsubsBtn-${props.postId}"`).style.display = "none";
            }).catch(error => {
                if (error.response && error.response.status === 401) {
                    // The server responded with a 400 status code
                    alert('please login again');
                    navigate("/login");
                }
                else{
                    console.log(error)
                    alert("Error occured please try again later.")
                }
            });
    }
    return (
        <div className="one-post">
            <h2>{props.title}</h2>
            <div className="bottom-section">
                <button 
                id={`"subsBtn-${props.postId}"`} 
                className="commentBtn" 
                onClick={Subscribe}
                style={
                    props.userSubscriptions.includes(props.postId)? {"display" : "none"} : {"display" : "block"}
                }
                >
                    Subscribe
                </button>
                <button 
                id={`"unsubsBtn-${props.postId}"`} 
                className="commentBtn" 
                onClick={Unsubscribe}
                style={
                    props.userSubscriptions.includes(props.postId)? 
                    {"backgroundColor": "red", "display" : "block"} : {"backgroundColor": "red", "display" : "none"}
                }
                >
                    Unsubscribe
                </button>
                <button className="commentBtn" onClick={openCommentPage}>Comments</button>
            </div>
        </div>
        
    )
}