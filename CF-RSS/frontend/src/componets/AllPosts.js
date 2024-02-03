import React from 'react';
import {useEffect, useState} from 'react';
const blogBox = document.getElementById("blogBox");

const PostSection = () => {
    const [isLoading, setIsLoading] = useState(false);
    
    useEffect(() => {
        const fetchPost = async () => {
            fetch('https://reactnative.dev/movies.json').then(response => response.json())
                .then(json => {
                    return json;
                }).catch(error => {
                console.error(error);
            });
        }
        console.log(fetchPost());
    }, []);
    if(isLoading) {
        return (
            <div>Loading</div>
        );
    }
    return  (
        <div id="#bolgBox">
        </div>
    );
};

export default PostSection;
