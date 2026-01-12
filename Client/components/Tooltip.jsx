'use client';

import { useEffect, useState } from 'react';
import ReactTooltip from 'react-tooltip';

// # 'Tooltip' Component #
const Tooltip = ({ id, place, offset, text }) => {
    const [theme, setTheme] = useState('light');

    useEffect(() => {
        // # Get the initial 'theme'
        const savedTheme = localStorage.getItem('theme') || 'light';
        setTheme(savedTheme);

        // # Watch for the 'theme' changes
        const observer = new MutationObserver(() => {
            const isDark = document.documentElement.classList.contains('dark');
            setTheme(isDark ? 'dark' : 'light');
        });

        observer.observe(document.documentElement, {
            attributes: true,
            attributeFilter: ['class']
        });

        return () => observer.disconnect();
    }, []);

    return (
        <>
            <ReactTooltip
                id={id}
                effect="solid"
                place={place}
                offset={offset}
                arrowColor={theme === 'light' ? '#e8eaed' : '#e7e8ec'}
                className={`tooltip ${theme === 'light' ? 'light-mode-tooltip' : 'dark-mode-tooltip'}`}
            >
                {text}
            </ReactTooltip>
        </>
    );
};

export default Tooltip;