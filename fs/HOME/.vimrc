set tabstop=4
set softtabstop=4
set shiftwidth=4
set expandtab
set autoindent
set backspace=indent,eol,start
autocmd FileType yaml setlocal ts=2 sts=2 sw=2 expandtab
syntax on
set number

if (has("termguicolors"))
    set termguicolors
endif
colorscheme onedark
