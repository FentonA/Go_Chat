import React from 'react';
import styles from './bubble.module.css'

const Bubble: React.FC = (props) =>{
    return  (
    <div className={styles.body}>
    <div className={styles.bubble}>
        <p className={styles.text}>
        This is a prop
        </p>
    </div>
        <p className={styles.time}>
            12:30pm
        </p>
    </div>
    
    )
}

export default Bubble;