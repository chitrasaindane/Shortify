'use client';

import { useRef, useState, useEffect, useMemo } from 'react';
import { Trash2, Copy, Check, ExternalLink } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import DeleteLinkDialog from './DeleteLinkDialog';
import Tooltip from './Tooltip';
import { getServerBaseURL } from '@/lib/url';

// # 'Link Card' Component #
const LinkCard = ({ link, username, onDelete, onClick, onCopy }) => {
    const timeoutRef = useRef(null);

    const [isDeleteOpen, setIsDeleteOpen] = useState(false);
    const [copied, setCopied] = useState(false);

    const linkShortUrl = useMemo(() => {
        const SERVER_BASE_URL = getServerBaseURL();
        const linkShortUrl = `${SERVER_BASE_URL}/${username}/${link.slug}`;
        return linkShortUrl;
    }, [username, link.slug]);

    useEffect(() => {
        // # Clear the 'timer' on the component 'unmount'
        return () => clearTimeout(timeoutRef.current);
    }, []);

    const handleCopy = (e) => {
        e.stopPropagation();
        onCopy(linkShortUrl);
        setCopied(true);

        const timer = setTimeout(() => {
            setCopied(false);
        }, 2000);

        // # Store the 'timer' in the 'ref'
        timeoutRef.current = timer;
    };

    return (
        <>
            <Card
                className="hover:shadow-lg transition-shadow cursor-pointer bg-gray-50 dark:bg-gray-800 dark:border-gray-700"
                onClick={() => onClick(link)}
            >
                <CardHeader>
                    <CardTitle className="flex items-start justify-between">
                        <span className="text-lg truncate pr-2 dark:text-white">
                            {link.title || 'Untitled'}
                        </span>
                        <Button
                            data-tip
                            data-for="delete-link"
                            variant="ghost"
                            size="icon"
                            className="h-9 w-9 hover:bg-gray-200 dark:hover:bg-gray-700"
                            onClick={(e) => {
                                e.stopPropagation();
                                setIsDeleteOpen(true);
                            }}
                        >
                            <Trash2 className="w-6 h-6 text-red-600" />
                        </Button>

                        {/* # 'Delete Link' Tooltip # */}
                        <Tooltip id="delete-link" place="bottom" offset={{ bottom: 5 }} text="Delete Link" />
                    </CardTitle>
                </CardHeader>
                <CardContent>
                    <div className="space-y-2">
                        <div>
                            <p className="text-xs text-gray-500 dark:text-gray-400 mb-1"> Short URL </p>
                            <div className="flex items-center space-x-2">
                                <code className="text-xs bg-gray-100 dark:bg-gray-700 px-1 py-1 flex-1 rounded truncate dark:text-gray-200">
                                    {linkShortUrl}
                                </code>
                                <Button
                                    data-tip
                                    data-for="copy-link"
                                    variant="ghost"
                                    size="icon"
                                    className={`h-9 w-9 hover:bg-gray-200 dark:hover:bg-gray-700 ${copied ? 'bg-gray-200 dark:bg-gray-700 animate-bounce' : ''}`}
                                    onClick={handleCopy}
                                >
                                    {copied ? (
                                        <Check className="w-6 h-6 text-green-600" />
                                    ) : (
                                        <Copy className="w-6 h-6" />
                                    )}
                                </Button>

                                {/* # 'Copy Link' Tooltip # */}
                                <Tooltip id="copy-link" place="bottom" offset={{ bottom: 5 }} text="Copy Link" />
                            </div>
                        </div>
                        <div>
                            <p className="text-xs text-gray-500 dark:text-gray-400 mb-1"> Destination URL </p>
                            <p className="text-xs text-gray-700 dark:text-gray-300 truncate"> {link.destination_url} </p>
                        </div>
                        <div className="flex justify-between items-center pt-2">
                            <span className="text-xs text-gray-500 dark:text-gray-200">
                                Slug:
                                <code className="bg-gray-100 dark:bg-gray-700 ml-1 px-1 py-1 flex-1 rounded truncate">
                                    {link.slug}
                                </code>
                            </span>
                            <a
                                data-tip
                                data-for="visit-link"
                                href={linkShortUrl}
                                target="_blank"
                                rel="noopener noreferrer"
                                onClick={(e) => e.stopPropagation()}
                                className="text-xs text-black dark:text-white hover:underline flex items-center"
                            >
                                Visit
                                <ExternalLink className="w-3 h-3 ml-1" />
                            </a>

                            {/* # 'Visit Link' Tooltip # */}
                            <Tooltip id="visit-link" place="bottom" offset={{ bottom: 5 }} text="Visit Link" />
                        </div>
                    </div>
                </CardContent>
            </Card >

            {/* # 'Delete Link' Dialog Component # */}
            <DeleteLinkDialog
                link={{ ...link, short_url: linkShortUrl }}
                open={isDeleteOpen}
                onOpenChange={setIsDeleteOpen}
                onConfirmDelete={onDelete}
            />
        </>
    );
};

export default LinkCard;