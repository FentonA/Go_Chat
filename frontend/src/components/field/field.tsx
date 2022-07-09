import React, {useRef} from 'react';
import Bubble from "../bubble/bubble"
import styles from "./field.module.css"
import LiveGollection from 'livegollection-client';

function generateClienttag(){
    return Math.random().toString(36).substring(2, 6)
}
const clientTag = generateClienttag()
const me = `Client#${clientTag}`

const Field: React.FC = (prop: any) =>{
    const liveGoll = new LiveGollection('ws://localhost:8080/livegollection')
    const messageInputRef = useRef<HTMLInputElement>(null);

     const submitHandler: React.FormEventHandler<HTMLFormElement> = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const enteredMessage = messageInputRef.current?.value
        addMessageToInbox(enteredMessage)

    }


    const addMessageToInbox = (message: any)=>{
        const sentByme = me == message.sender;

        liveGoll.create({
            sender: me,
            sentTime: new Date(),
            text: message
        })

    }
    return (
        <div className={styles.body}>
            <form onSubmit={submitHandler} className={styles.input}> 
                <input
                    type="text" 
                    placeholder="Send a message"
                    required ref={messageInputRef}
                    >
                </input>
                <button type='submit'>Send</button>
            </form>
        </div>
    )
}

export default  Field