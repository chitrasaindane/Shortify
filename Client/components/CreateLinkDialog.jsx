'use client';

import { useEffect, useState } from 'react';
import { Plus, Loader2 } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import Toast from './Toast';

// # 'Create Link' Dialog Component #
const CreateLinkDialog = ({ open, onOpenChange, onCreate }) => {
    const [formData, setFormData] = useState({
        title: '',
        description: '',
        destination_url: '',
        slug: '',
    });

    const [toast, setToast] = useState(null);
    const [isLoading, setIsLoading] = useState(false);

    // # Clear the 'form' on the 'dialog' close
    useEffect(() => {
        if (!open) {
            setFormData({ title: '', description: '', destination_url: '', slug: '' });
            setIsLoading(false);
        }
    }, [open]);

    const showToast = (message, type) => {
        setToast({ message, type });
    };

    const handleSubmit = async () => {
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
        await onCreate(formData);
        setIsLoading(false);
    };

    return (
        <>
            <Dialog open={open} onOpenChange={onOpenChange}>
                <DialogTrigger asChild>
                    <Button className="bg-black hover:bg-gray-800 dark:bg-white dark:text-black dark:hover:bg-gray-200">
                        <Plus className="w-4 h-4" />
                        Create Link
                    </Button>
                </DialogTrigger>
                <DialogContent className="dark:bg-gray-800 dark:border-gray-700">
                    <DialogHeader>
                        <DialogTitle className="dark:text-white">
                            Create New Link
                        </DialogTitle>
                    </DialogHeader>
                    <div className="space-y-4 mt-4">
                        <div>
                            <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                Title (optional)
                            </label>
                            <Input
                                value={formData.title}
                                onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                                placeholder="My Link"
                                disabled={isLoading}
                                className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                            />
                        </div>
                        <div>
                            <label className="text-sm font-medium mb-1 block dark:text-gray-200">
                                Description (optional)
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
                                placeholder="my-custom-slug"
                                disabled={isLoading}
                                className="bg-gray-100 dark:bg-gray-700 dark:text-white dark:border-gray-600"
                            />
                        </div>
                        <Button
                            onClick={handleSubmit}
                            disabled={isLoading}
                            className="w-full bg-black hover:bg-gray-800 dark:bg-white dark:text-black dark:hover:bg-gray-200 disabled:opacity-50"
                        >
                            {isLoading ? (
                                <>
                                    <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                                </>
                            ) : (
                                'Create Link'
                            )}
                        </Button>
                    </div>
                </DialogContent>
            </Dialog>

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

export default CreateLinkDialog;