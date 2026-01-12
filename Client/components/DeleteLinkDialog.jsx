'use client';

import { useState } from 'react';
import { Trash2, AlertTriangle, Loader2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription } from '@/components/ui/dialog';

// # 'Delete Link' Dialog Component #
const DeleteLinkDialog = ({ link, open, onOpenChange, onConfirmDelete }) => {
    const [inputSlug, setInputSlug] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleOpenChange = (isOpen) => {
        if (!isOpen) {
            setInputSlug('');
            setIsLoading(false);
        }
        onOpenChange(isOpen);
    };

    const handleDelete = async () => {
        setIsLoading(true);
        await onConfirmDelete(link._id);
        setIsLoading(false);
        handleOpenChange(false);
    };

    const isSlugMatch = inputSlug === link?.slug;

    if (!link) return null;

    return (
        <>
            <Dialog open={open} onOpenChange={handleOpenChange}>
                <DialogContent className="dark:bg-gray-800 dark:border-gray-700 max-w-xl">
                    <DialogHeader>
                        <DialogTitle className="flex items-center space-x-2 text-red-600 dark:text-red-400">
                            <AlertTriangle className="w-5 h-5" />
                            <span> Delete Link </span>
                        </DialogTitle>
                        <DialogDescription className="dark:text-gray-300">
                            This action cannot be undone. Please confirm by entering the link slug.
                        </DialogDescription>
                    </DialogHeader>

                    <div className="space-y-4 mt-4">
                        {/* # 'Link' Information #*/}
                        <div className="bg-gray-100 dark:bg-gray-900 p-4 rounded-lg space-y-3">
                            <div>
                                <p className="text-xs text-gray-500 dark:text-gray-400 mb-1"> Title </p>
                                <p className="text-sm font-medium text-gray-900 dark:text-white">
                                    {link.title || 'Untitled'}
                                </p>
                            </div>

                            <div>
                                <p className="text-xs text-gray-500 dark:text-gray-400 mb-1"> Destination URL </p>
                                <p className="text-sm text-gray-900 dark:text-white break-all">
                                    {link.destination_url}
                                </p>
                            </div>

                            <div>
                                <p className="text-xs text-gray-500 dark:text-gray-400 mb-1"> Short URL </p>
                                <code className="text-sm text-gray-900 dark:text-white break-all">
                                    {link.short_url}
                                </code>
                            </div>

                            <div>
                                <p className="text-xs text-gray-500 dark:text-gray-400 mb-1"> Slug to confirm </p>
                                <code className="text-sm font-bold text-red-600 dark:text-red-400">
                                    {link.slug}
                                </code>
                            </div>
                        </div>

                        {/* # 'Input' Field # */}
                        <div>
                            <label className="text-sm font-medium mb-2 block dark:text-gray-200">
                                Type
                                <span className="font-bold text-red-600 dark:text-red-400"> {link.slug} </span>
                                to confirm
                            </label>
                            <Input
                                value={inputSlug}
                                onChange={(e) => setInputSlug(e.target.value)}
                                placeholder="Enter slug here"
                                disabled={isLoading}
                                className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                                autoFocus
                            />
                        </div>

                        {/* # 'Action' Buttons # */}
                        <div className="flex space-x-2 pt-2">
                            <Button
                                onClick={() => handleOpenChange(false)}
                                disabled={isLoading}
                                className="flex-1 bg-black hover:bg-gray-800 dark:bg-white dark:text-black dark:hover:bg-gray-200 disabled:opacity-50"
                            >
                                Cancel
                            </Button>
                            <Button
                                onClick={handleDelete}
                                disabled={!isSlugMatch || isLoading}
                                className="flex-1 bg-red-600 hover:bg-red-700 text-white disabled:opacity-50 disabled:cursor-not-allowed"
                            >
                                {isLoading ? (
                                    <>
                                        <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                                    </>
                                ) : (
                                    <>
                                        <Trash2 className="w-4 h-4 mr-2" />
                                        Confirm Delete
                                    </>
                                )}
                            </Button>
                        </div>
                    </div>
                </DialogContent>
            </Dialog>
        </>
    );
};

export default DeleteLinkDialog;