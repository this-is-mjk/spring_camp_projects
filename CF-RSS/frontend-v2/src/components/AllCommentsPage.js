import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import "./AllCommentsPage.css"
import { IoIosContact } from "react-icons/io";
export default function AllCommentsPage () {
    const location = useLocation();
    const {postId} = location.state;
    const [actions, setActions] = useState([]);
    
    const timeConverter = (unixTimestamp) => {
        const date = new Date(unixTimestamp * 1000); // Convert Unix timestamp to milliseconds
        // Get the various components of the date
        const year = date.getFullYear();
        const month = date.getMonth() + 1; // Month is 0-indexed, so we add 1
        const day = date.getDate();
        const hours = date.getHours();
        const minutes = date.getMinutes();
        const seconds = date.getSeconds();
        return(`${year}-${month}-${day} ${hours}:${minutes}:${seconds}`)
    }
    
    useEffect(() => {
        // here json.parse will give a object with key value pairs
        const posts = JSON.parse(sessionStorage.getItem('groupedCommentData'));
        // console.log(posts)
        if (posts && posts[postId]) {
            const post = posts[postId];
            setActions(post);
        } else {
            console.error(`Post with ID ${postId} not found in posts`);
        }
    }, [postId]);
    return (
        <div className="commentsection">
            <div className="bgimageCommentSection"></div>
             {/* <h1>{postId}</h1> */}
             {actions.length > 0 && (
                <div>
                    <div className="post-details">
                        <div className="time-author">
                            <div style={{display:"flex", alignItems:"center"}}>
                                <IoIosContact />
                                <p>{actions[0].blogEntry.authorHandle}</p>
                            </div>
                            <p>{timeConverter(actions[0].blogEntry.creationTimeSeconds)}</p>
                        </div>
                        <h1 className="heading">{actions[0].blogEntry.title.slice(3, -4)}</h1>
                    </div>
                    <div className="comment-top-layer">
                    {actions.map((action, index) => (
                        <div key={index}>
                            {action.comment && (
                                <div className="comment" >
                                    <div className="time-author">
                                        <div style={{display:"flex", alignItems:"center"}}>
                                            <IoIosContact />
                                            <p>{action.comment.commentatorHandle}</p>
                                        </div>  
                                        <p>{timeConverter(action.comment.creationTimeSeconds)}</p>
                                    </div>
                                    <div dangerouslySetInnerHTML={{ __html: action.comment.text }} />
                                </div>
                        )}
                        </div>
                    ))}
                    </div>
                </div>
            )}
        </div>
       
    );
}