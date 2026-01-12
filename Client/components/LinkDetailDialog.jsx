'use client';

import { useRef, useState, useEffect, useMemo } from 'react';
import { Copy, Check, Trash2, Loader2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog';
import { Alert, AlertDescription } from '@/components/ui/alert';
import DeleteLinkDialog from './DeleteLinkDialog';
import Toast from './Toast';
import { getServerBaseURL } from '@/lib/url';

// # 'Link Detail Dialog' Component #
const LinkDetailDialog = ({ link, username, open, onOpenChange, onUpdate, onDelete }) => {
    const [formData, setFormData] = useState({
        title: '',
        description: '',
        destination_url: '',
        slug: '',
    });

    const [isDeleteOpen, setIsDeleteOpen] = useState(false);
    const [copied, setCopied] = useState(false);
    const [toast, setToast] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    const inputRef = useRef(null);
    const timeoutRef = useRef(null);

    const linkShortUrl = useMemo(() => {
        if (!link) {
            return '';
        }
        const SERVER_BASE_URL = getServerBaseURL();
        const linkShortUrl = `${SERVER_BASE_URL}/${username}/${link.slug}`;
        return linkShortUrl;
    }, [username, link]);

    useEffect(() => {
        if (link) {
            setFormData({
                title: link.title || '',
                description: link.description || '',
                destination_url: link.destination_url || '',
                slug: link.slug || '',
            });

            // # Prevent the 'auto-selection' of the 'title' (input text) by setting the 'cursor' at the 'end' after the 'render'
            requestAnimationFrame(() => {
                if (inputRef.current) {
                    const len = (link.title || '').length;
                    inputRef.current.setSelectionRange(len, len);
                    inputRef.current.focus();
                }
            });
        }
    }, [link]);

    useEffect(() => {
        // # Clear the 'timer' on the component 'unmount'
        return () => clearTimeout(timeoutRef.current);
    }, []);

    useEffect(() => {
        if (!open) {
            setIsLoading(false);
        }
    }, [open]);

    const handleCopy = () => {
        copyToClipboard(linkShortUrl);
        setCopied(true);

        const timer = setTimeout(() => {
            setCopied(false);
        }, 2000);

        // # Store the 'timer' in the 'ref'
        timeoutRef.current = timer;
    };

    const showToast = (message, type) => {
        setToast({ message, type });
    };

    const handleUpdate = async () => {
        if (
            !formData.title.trim() &&
            !formData.description.trim() &&
            !formData.destination_url.trim() &&
            !formData.slug.trim()
        ) {
            // # Close the 'dialog' and show the 'toast'
            const outerTimer = setTimeout(() => {
                // # Close the 'dialog' after '500ms'
                onOpenChange(false);

                // # Show the 'toast' after another '500ms'
                innerTimer = setTimeout(() => {
                    const errMsg = 'Please fill out the form first';
                    showToast(errMsg, 'error');
                }, 500);
            }, 500);

            let innerTimer; // # Declare the 'inner timer' variable 'outer' so that 'cleanup' can see it

            // # Clear the 'timer' (inner & outer) on the component 'unmount'
            return () => {
                clearTimeout(innerTimer);
                clearTimeout(outerTimer);
            };
        }

        setIsLoading(true);
        await onUpdate(link._id, formData);
        setIsLoading(false);
    };

    // # Copy To 'Clipboard'
    const copyToClipboard = (text) => {
        navigator.clipboard.writeText(text);
    };

    if (!link) return null;

    return (
        <>
            <Dialog open={open} onOpenChange={onOpenChange}>
                <DialogContent className="max-w-2xl dark:bg-gray-800 dark:border-gray-700">
                    <DialogHeader>
                        <DialogTitle className="dark:text-white">
                            Link Details
                        </DialogTitle>
                    </DialogHeader>
                    <div className="space-y-4 mt-4">
                        <div>
                            <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                Title
                            </label>
                            <Input
                                ref={inputRef}
                                value={formData.title}
                                onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                                placeholder="My Link"
                                disabled={isLoading}
                                className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                            />
                        </div>

                        <div>
                            <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                Description
                            </label>
                            <Textarea
                                value={formData.description}
                                onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                                placeholder="Link Description"
                                disabled={isLoading}
                                className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                            />
                        </div>

                        <div>
                            <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                Destination URL <span className="text-red-500"> * </span>
                            </label>
                            <Input
                                value={formData.destination_url}
                                onChange={(e) => setFormData({ ...formData, destination_url: e.target.value })}
                                placeholder="https://example.com"
                                disabled={isLoading}
                                className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                            />
                        </div>

                        <div>
                            <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                Slug <span className="text-red-500"> * </span>
                            </label>
                            <Input
                                value={formData.slug}
                                onChange={(e) => setFormData({ ...formData, slug: e.target.value.toLowerCase().trim() })}
                                placeholder="custom-slug"
                                disabled={isLoading}
                                className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                            />
                        </div>

                        {copied && (
                            <Alert className="bg-green-100 border border-green-500">
                                <AlertDescription className="text-md font-medium text-green-700">
                                    Link copied to clipboard!
                                </AlertDescription>
                            </Alert>
                        )}

                        <div>
                            <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                Short URL
                            </label>
                            <div className="flex items-center space-x-2">
                                <Input
                                    value={linkShortUrl}
                                    readOnly
                                    className="flex-1 bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                                />
                                <Button
                                    onClick={handleCopy}
                                    disabled={isLoading}
                                    className={`bg-black hover:bg-gray-800 dark:bg-white dark:text-black dark:hover:bg-gray-200 ${copied ? 'animate-bounce' : ''}`}
                                >
                                    {copied ? (
                                        <Check className="w-6 h-6 text-green-600" />
                                    ) : (
                                        <Copy className="w-6 h-6" />
                                    )}
                                </Button>
                            </div>
                        </div>

                        <div className="grid grid-cols-2 gap-4">
                            <div>
                                <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                    Created At
                                </label>
                                <Input
                                    value={
                                        new Date(link.created_at).toLocaleString(
                                            'en-IN',
                                            {
                                                day: '2-digit',
                                                month: 'long',
                                                year: 'numeric',
                                                hour: '2-digit',
                                                minute: '2-digit',
                                                second: '2-digit',
                                                hour12: true
                                            }
                                        )
                                            .replace(' at ', ' ')
                                            .replace('am', 'AM')
                                            .replace('pm', 'PM')
                                    }
                                    readOnly
                                    className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                                />
                            </div>

                            {link.updated_at && (
                                <div>
                                    <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                        Updated At
                                    </label>
                                    <Input
                                        value={
                                            new Date(link.updated_at).toLocaleString(
                                                'en-IN',
                                                {
                                                    day: '2-digit',
                                                    month: 'long',
                                                    year: 'numeric',
                                                    hour: '2-digit',
                                                    minute: '2-digit',
                                                    second: '2-digit',
                                                    hour12: true
                                                }
                                            )
                                                .replace(' at ', ' ')
                                                .replace('am', 'AM')
                                                .replace('pm', 'PM')
                                        }
                                        readOnly
                                        className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                                    />
                                </div>
                            )}
                        </div>

                        <div className="flex space-x-2 pt-4">
                            <Button
                                onClick={handleUpdate}
                                disabled={isLoading}
                                className="flex-1 bg-black hover:bg-gray-800 dark:bg-white dark:text-black dark:hover:bg-gray-200 disabled:opacity-50"
                            >
                                {isLoading ? (
                                    <>
                                        <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                                    </>
                                ) : (
                                    'Update Link'
                                )}
                            </Button>
                            <Button
                                onClick={() => setIsDeleteOpen(true)}
                                disabled={isLoading}
                                className="bg-red-600 hover:bg-red-700 text-white disabled:opacity-50"
                            >
                                <Trash2 className="w-4 h-4" />
                                Delete Link
                            </Button>
                        </div>
                    </div>
                </DialogContent>
            </Dialog>

            {/* # 'Delete Link' Dialog Component # */}
            <DeleteLinkDialog
                link={{ ...link, short_url: linkShortUrl }}
                open={isDeleteOpen}
                onOpenChange={setIsDeleteOpen}
                onConfirmDelete={(id) => {
                    onDelete(id);
                    onOpenChange(false);
                }}
            />

            {/* # 'Toast' Component # */}
            {toast && (
                <Toast
                    message={toast.message}
                    type={toast.type}
                    onClose={() => setToast(null)}
                />
            )}
        </>
    );
};

export default LinkDetailDialog;