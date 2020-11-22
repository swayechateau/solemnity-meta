<?php
namespace App\Helpers;

class Helper
{
    // Grab meta data from site
    public static function getMetaData($link, $all = false)
    {
        $all = Helper::toBool($all);
        // site user agent to firefox, to minimise site user agent restrictions
        $userAgent = 'Mozilla/5.0 (Windows NT 5.1; rv:31.0) Gecko/20100101 Firefox/31.0';
        ini_set('user_agent', $userAgent);

        // check if link has http else add it
        if (strpos($link, 'http') === false) {
            $link = 'http://' . $link;
        }
        
        // create a new html document element, to spoof viewing as a client
        $html = new \DOMDocument();
        // prepare the site contents to be scrapped
        $sites_html = file_get_contents($link);
        @$html->loadHTML($sites_html);

        // get response array ready
        $response = [];
        //Get all meta tags and loop through them.
        $response['title'] = $html->getElementsByTagName('title')[0]->textContent;
        $response['url'] = $link;
        if ($all === true) {
            $response['meta'] = [];
        }
        foreach ($html->getElementsByTagName('meta') as $meta) {
            if ($all === true) {
                if($meta->getAttribute('name')) {
                    $response['meta'][$meta->getAttribute('name')] = $meta->getAttribute('content');
                }else if($meta->getAttribute('rel')) {
                    $response['meta'][$meta->getAttribute('rel')] = $meta->getAttribute('content');
                }else if($meta->getAttribute('itemprop')) {
                    $response['meta'][$meta->getAttribute('itemprop')] = $meta->getAttribute('content');
                }else if($meta->getAttribute('property')) {
                    $response['meta'][$meta->getAttribute('property')] = $meta->getAttribute('content');
                }
            } else {
                if (Helper::has_meta_description($meta)) {
                    $response['description'] = $meta->getAttribute('content');
                };

                if (Helper::has_meta_image($meta)) {
                    $response['image'] = $meta->getAttribute('content');
                }

                if (Helper::has_meta_url($meta)) {
                    $response['url'] = $meta->getAttribute('content');
                }
                if ($meta->getAttribute('name') === 'twitter:domain') {
                    return $meta->getAttribute('content');
                }
            }
        }
        return $response;
    }

    public static function has_meta_title($meta)
    {
        if ($meta->getAttribute('itemprop') == 'title' 
            || $meta->getAttribute('property') == 'og:title'
            || $meta->getAttribute('name') == 'twitter:title'
            || $meta->getAttribute('name') == 'title') {
            return true;
        }
        return false;
    }

    public static function has_meta_image($meta)
    {
        if ($meta->getAttribute('itemprop') == 'image' 
            || $meta->getAttribute('property') == 'og:image' 
            || $meta->getAttribute('name') == 'twitter:image' 
            || $meta->getAttribute('name') == 'twitter:image:src'
            || $meta->getAttribute('name') == 'image') {
            return true;
        }
        return false;
    }

    public static function has_meta_description($meta)
    {
        if ($meta->getAttribute('itemprop') == 'description'
            || $meta->getAttribute('property') == 'og:description' 
            || $meta->getAttribute('name') == 'twitter:description' 
            || $meta->getAttribute('name') == 'description') {
            return true;
        }
        return false;
    }

    public static function has_meta_url($meta)
    {
        if ($meta->getAttribute('itemprop') == 'url' 
            || $meta->getAttribute('property') == 'og:url'
            || $meta->getAttribute('name') == 'twitter:url'
            || $meta->getAttribute('name') == 'url') {
            return true;
        }
        return false;
    }

    // convert boolean
    public static function toBool($bool) {
        if($bool === true || $bool === 1 || $bool === 'true' || $bool === '1') {
            return true;
        }
        return false;
    }
}
