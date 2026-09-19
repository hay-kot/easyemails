// Package easyemails provides a simple abstraction around common email templates
// for sending transactional emails for things like
//
//   - Welcome emails
//   - Password reset emails
//   - Email verification emails
//   - Magic Links
//   - etc
//
// Functions with a With prefix, such as WithParagraph and WithButton, create
// the components of an email. The prefix keeps components apart from the
// other functions and methods in the package.
package easyemails
