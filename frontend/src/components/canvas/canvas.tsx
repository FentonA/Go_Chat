import React from 'react';
import Bubble from "../bubble/bubble"
import Field from "../field/field"
import styles from './canvas.module.css'
import LiveGollection from 'livegollection-client';

const Canvas: React.FC = () =>{
    return (
        <div className={styles.body}>
            <Bubble/>
            <Field/>
        </div>
    )
}

export default Canvas