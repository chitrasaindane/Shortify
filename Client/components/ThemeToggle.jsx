'use client';

import { useEffect, useState } from 'react';
import { Moon, Sun } from 'lucide-react';
import { Button } from '@/components/ui/button';
import Tooltip from './Tooltip';

// # 'Theme Toggle' Component #
const ThemeToggle = () => {
    const [theme, setTheme] = useState('light');

    useEffect(() => {
        const savedTheme = localStorage.getItem('theme') || 'light';
        setTheme(savedTheme);
        document.documentElement.classList.toggle('dark', savedTheme === 'dark');
    }, []);

    const toggleTheme = () => {
        const newTheme = theme === 'light' ? 'dark' : 'light';
        setTheme(newTheme);
        localStorage.setItem('theme', newTheme);
        document.documentElement.classList.toggle('dark', newTheme === 'dark');
    };

    return (
        <>
            <Button data-tip data-for="theme-toggle" variant="ghost" size="icon" className="h-9 w-9" onClick={toggleTheme}>
                {theme === 'light' ? (
                    <Moon className="w-6 h-6" />
                ) : (
                    <Sun className="w-6 h-6" />
                )}
            </Button>

            {/* # 'Theme Toggle' Tooltip # */}
            <Tooltip id="theme-toggle" place="bottom" offset={{ bottom: 5 }} text={theme === 'light' ? 'Switch to Dark Mode' : 'Switch to Light Mode'} />
        </>
    );
};

export default ThemeToggle;
