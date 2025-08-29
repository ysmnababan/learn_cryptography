# CSRF

- CSRF: Cross-Site Request Forgery  
- Is an attack that uses the valid credentials of a user, and acts like that particular user
  accessing a certain web. This can be harmful because the process can be done 
  without the user noticing. This happens because the attack is done by exploiting the user's cookies.  

## Cookies
- Cookies are a mechanism used for storing user data.  
- By using these cookies, every request that the user sends can be validated by the server because 
  the request itself contains user information.  
- The cookies normally store the session id of a client that can be checked if it is a valid session
  when it arrives at the server.  
- Cookies are natively supported in the browser and are sent automatically by the browser for each request.  
  The client doesn't have to set it up in the request manually, because it is already handled by the browser by 
  setting it directly in the request header.  

  ```
  Set-Cookie: sessionId=abc123; Domain=example.com; Path=/; Secure; HttpOnly; SameSite=None
  ```

- This cookie will be sent for `example.com` and its subdomains (`api.example.com`, `www.example.com`).  
- The **Domain flag** is a way of telling the browser: *if you send a request to this specified domain, attach the cookie.*  
  So in that example, the browser will automatically attach the cookie when sending a request to `example.com` and 
  its subdomains like `api.example.com`, but not to a different target like `anotherexample.com`.  
- The cookie is valid for all URLs under `/`.  
- Another important flag of cookies is `SameSite`.  
- This is for specifying when the browser is allowed to include cookies in cross-site requests.  

## How CSRF happens
- `SameSite=None` means cookies will be sent with cross-site requests, so for example `attacker.com` can create a request 
  to `example.com` and it will be recognized as if coming from a valid user.  
- The flow: you log in and the browser stores your cookies. You then open a malicious site, and it creates a request 
  to `example.com`. The browser already has the cookies, so it will send them automatically, 
  because the cookies have the `None` configuration, allowing any site to trigger requests with them.  
- To prevent this, you can use another option for that flag, `SameSite=Strict` or `Lax`. This mechanism makes sure
  that requests only come from the same site. So `attacker.com` can't make the browser send a valid request to `example.com`.  
- You can also add the `Secure` flag to make sure cookies are only sent over HTTPS.  
- By configuring `SameSite`, it becomes difficult to support multiple non-browser clients like mobile apps.  
  This is because the client is considered a different site, and its requests won't automatically carry cookies, 
  even if it is your own mobile app. This is why cookies are more commonly used for SPA pages where the FE and BE are typically
  hosted on the same domain.  
- When it comes to multi-device apps, it is common to use something like a JWT (stored
  in client localStorage or sessionStorage) because it is not tied
  to browser cookie behavior and there is no auto-attach like cookies.  
- But this approach also has some flaws, mainly XSS attacks — because malicious scripts can access storage using JS
  and steal the token. This can be tackled with other approaches that are not in scope here.  

**Trade-off:**  
- Cookies = easier for browsers, CSRF risk (mitigated by `SameSite`)  
- JWT = easier for multi-device APIs, CSRF-safe, but XSS risk