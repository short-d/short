import React from 'react';
import './App.scss';

function App() {
    return (
        <div className='App'>
            <header>
                <div className={'center'}>
                    <div className={'logo'}>Short</div>
                </div>
            </header>
            <div className={'content'}>
                <div className={'center'}>

                    <div className={'title'}>New Short Link</div>

                    <div className={'control create-short-link'}>
                        <input className={'text-field'} type={'text'} placeholder={'Long Link'}/>
                        <input className={'text-field'} type={'text'} placeholder={'Custom Short Link ( Optional )'}/>
                        <button>
                            Create Short Link
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default App;
