const gulp = require('gulp');
const markdownPdf = require('gulp-markdown-pdf');
var del = require('del');

del('./*.pdf');   
 
gulp.task('default', () =>
    gulp.src('*.md')
        .pipe(markdownPdf())
        .pipe(gulp.dest('dist'))
);