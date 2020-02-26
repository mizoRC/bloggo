import React from 'react';
import ReactMarkdown from 'react-markdown';
import MDEditor, { commands } from '@uiw/react-md-editor';

const Test = () => {
    const [ post, setPost ] = React.useState('**Hola**'); 

    const changePost = value => {
        setPost(value);
    }

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
                <MDEditor
                    value={post}
                    onChange={changePost}
                    preview={'edit'}
                    commands={[
                        commands.bold, commands.italic, commands.strikethrough, commands.title, commands.hr, commands.divider, commands.link, commands.quote, commands.code, commands.image, commands.unorderedListCommand, commands.orderedListCommand, commands.checkedListCommand, commands.divider,
                        commands.fullscreen
                    ]}
                    visiableDragbar={false}
                    style={{width: '100%', height: '100%'}}
                />
            </div>

            <div
                style={{
                    width: '50%',
                    height: '100%',
                    padding: '10px'
                }}
            >
                <ReactMarkdown source={post} />
            </div>
        </div>
    )
}

export default Test;