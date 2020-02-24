import React from 'react';
import ReactMarkdown from 'react-markdown';

const Test = () => {
    const [ post, setPost ] = React.useState(''); 

    return (
        <div
            style={{
                width: '100%',
                height: '100%',
                display: 'flex',
                flexDirection: 'row',
                alignItems: 'center',
                justifyContent: 'center'
            }}
        >
            <div
                style={{
                    width: '50%',
                    height: '100%',
                    display: 'flex',
                    flexDirection: 'row',
                    alignItems: 'center',
                    justifyContent: 'center'
                }}
            >
                <ReactMarkdown source={post} />
            </div>

            <div
                style={{
                    width: '50%',
                    height: '100%',
                    display: 'flex',
                    flexDirection: 'row',
                    alignItems: 'center',
                    justifyContent: 'center'
                }}
            >
            
            </div>
        </div>
    )
}

export default Test;